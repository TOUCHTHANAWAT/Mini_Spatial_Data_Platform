import { Modal, Form, Input, Select, message } from "antd";
import { useEffect, useState } from "react";
import { createFeature, getCategories } from "../api/placeApi";
import { Category } from "../interfaces/place";

type Props = {
  open: boolean;
  onClose: () => void;
  onSuccess: () => void;
};

export default function CreatePlaceModal({
  open,
  onClose,
  onSuccess,
}: Props) {
  const [form] = Form.useForm();
  const [categories, setCategories] = useState<Category[]>([]);

  useEffect(() => {
    if (!open) return;

    const fetchCategories = async () => {
      try {
        const data = await getCategories();
        setCategories(data ?? []);
      } catch (err) {
        message.error("Category loading failed.");
      }
    };

    fetchCategories();
  }, [open]);

  const handleSubmit = async (values: any) => {
    try {
      await createFeature({
        type: "Feature",
        geometry: {
          type: "Point",
          coordinates: [Number(values.lng), Number(values.lat)],
        },
        properties: {
          name: values.name,
          category: values.category,
        },
      });

      message.success("Created successfully");

      form.resetFields();
      onClose();
      onSuccess();
    } catch (err) {
      message.error("Create failed");
      console.error(err);
    }
  };

  return (
    <Modal
      title="Create Place"
      open={open}
      onCancel={() => {
        form.resetFields();
        onClose();
      }}
      onOk={() => form.submit()}
      okText="Create"
      cancelText="Cancel"
    >
      <Form form={form} layout="vertical" onFinish={handleSubmit}>

        <Form.Item
          name="name"
          label="Name"
          rules={[
            { required: true, message: "Please enter place name" },
            { min: 2, message: "Name must be at least 2 characters" },
          ]}
        >
          <Input placeholder="e.g. Bangkok City Park" />
        </Form.Item>

        <Form.Item
          name="category"
          label="Category"
          rules={[
            { required: true, message: "Please select a category" },
          ]}
        >
          <Select placeholder="Select a category">
            {categories.map((c) => (
              <Select.Option key={c.value} value={c.value}>
                {c.label}
              </Select.Option>
            ))}
          </Select>
        </Form.Item>


        <Form.Item
          name="lng"
          label="Longitude"
          rules={[
            { required: true, message: "Longitude is required" },
            {
              validator: (_, value) => {
                if (!value) return Promise.resolve();

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
          <Input placeholder="e.g. 100.5231" />
        </Form.Item>

        <Form.Item
          name="lat"
          label="Latitude"
          rules={[
            { required: true, message: "Latitude is required" },
            {
              validator: (_, value) => {
                if (!value) return Promise.resolve();

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
          <Input placeholder="e.g. 13.7367" />
        </Form.Item>

      </Form>
    </Modal>
  );
}