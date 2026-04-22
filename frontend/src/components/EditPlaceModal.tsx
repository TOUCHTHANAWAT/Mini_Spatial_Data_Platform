import { useEffect } from "react";
import { Modal, Form, Input, Select, message } from "antd";
import { updateFeature } from "../api/placeApi";
import { Feature, Category } from "../interfaces/place";

type Props = {
  open: boolean;
  feature: Feature | null;
  categories: Category[];
  onClose: () => void;
  onSuccess: () => void;
};

export default function EditPlaceModal({
  open,
  feature,
  categories,
  onClose,
  onSuccess,
}: Props) {
  const [form] = Form.useForm();

  useEffect(() => {
    if (feature) {
      form.setFieldsValue({
        name: feature.properties.name,
        category: feature.properties.category,
        lng: feature.geometry.coordinates[0],
        lat: feature.geometry.coordinates[1],
      });
    } else {
      form.resetFields();
    }
  }, [feature, form]);

  const handleSubmit = async (values: any) => {
    if (!feature) return;

    try {
      await updateFeature(feature.id, {
        properties: {
          name: values.name,
          category: values.category,
        },
        geometry: {
          type: "Point",
          coordinates: [Number(values.lng), Number(values.lat)],
        },
      });

      message.success("Updated successfully");

      form.resetFields();
      onClose();
      onSuccess();
    } catch (err) {
      console.error(err);
      message.error("Update failed");
    }
  };

  return (
    <Modal
      title="Edit Feature"
      open={open}
      onCancel={() => {
        form.resetFields();
        onClose();
      }}
      onOk={() => form.submit()}
      destroyOnClose
    >
      <Form form={form} layout="vertical" onFinish={handleSubmit}>
        <Form.Item
          label="Name"
          name="name"
          rules={[{ required: true, message: "Please input name" }]}
        >
          <Input />
        </Form.Item>

        <Form.Item
          name="category"
          label="Category"
          rules={[{ required: true, message: "Please select category" }]}
        >
          <Select placeholder="Select category">
            {categories.map((c) => (
              <Select.Option key={c.value} value={c.value}>
                {c.label}
              </Select.Option>
            ))}
          </Select>
        </Form.Item>

        <Form.Item
          label="Longitude"
          name="lng"
          rules={[
            { required: true, message: "Please input longitude" },
            {
              validator: (_, value) => {
                if (value === undefined || value === null || value === "") {
                  return Promise.resolve();
                }

                const num = Number(value);
                if (isNaN(num)) {
                  return Promise.reject("Enter valid longitude number");
                }

                if (num < -180 || num > 180) {
                  return Promise.reject("Longitude must be between -180 and 180");
                }

                return Promise.resolve();
              },
            },
          ]}
        >
          <Input  />
        </Form.Item>
        <Form.Item
          label="Latitude"
          name="lat"
          rules={[
            { required: true, message: "Please input latitude" },
            {
              validator: (_, value) => {
                if (value === undefined || value === null || value === "") {
                  return Promise.resolve();
                }

                const num = Number(value);
                if (isNaN(num)) {
                  return Promise.reject("Enter valid latitude number");
                }

                if (num < -90 || num > 90) {
                  return Promise.reject("Latitude must be between -90 and 90");
                }

                return Promise.resolve();
              },
            },
          ]}
        >
          <Input />
        </Form.Item>
      </Form>
    </Modal>
  );
}