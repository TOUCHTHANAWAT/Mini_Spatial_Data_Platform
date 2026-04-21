import { Form, Input, Button } from "antd";
import { useNavigate } from "react-router-dom";
import { createFeature } from "../api/placeApi";

export default function CreatePlacePage() {
  const navigate = useNavigate();

  const onFinish = async (values: any) => {
    try {
      await createFeature({
        type: "Feature",
        geometry: {
          type: "Point",
          coordinates: [
            parseFloat(values.lng),
            parseFloat(values.lat),
          ],
        },
        properties: {
          name: values.name,
        },
      });

      navigate("/places");
    } catch (error) {
      console.error("Create failed:", error);
    }
  };

  return (
    <div style={{ maxWidth: 500, margin: "0 auto", padding: 20 }}>
      <h2>Create Place</h2>

      <Form onFinish={onFinish} layout="vertical">
        <Form.Item
          name="name"
          label="Name"
          rules={[{ required: true, message: "Name is required" }]}
        >
          <Input placeholder="e.g. Bangkok" />
        </Form.Item>

        <Form.Item
          name="lng"
          label="Longitude"
          rules={[{ required: true, message: "Longitude is required" }]}
        >
          <Input placeholder="e.g. 100.5231" />
        </Form.Item>

        <Form.Item
          name="lat"
          label="Latitude"
          rules={[{ required: true, message: "Latitude is required" }]}
        >
          <Input placeholder="e.g. 13.7367" />
        </Form.Item>

        <Button type="primary" htmlType="submit" block>
          Create Feature
        </Button>
      </Form>
    </div>
  );
}