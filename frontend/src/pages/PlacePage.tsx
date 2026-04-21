import { useEffect, useState } from "react";
import { Table, Button } from "antd";
import { getFeatures } from "../api/placeApi";

export default function PlacePage() {
  const [features, setFeatures] = useState<any[]>([]);

  const load = async () => {
    const data = await getFeatures();
    setFeatures(data);
  };

  useEffect(() => {
    load();
  }, []);

  return (
    <Table
      dataSource={features}
      rowKey="id"
      columns={[
        { title: "Name", dataIndex: ["properties", "name"] },

        {
          title: "Latitude",
          render: (f) => {
            const lat = f.geometry.coordinates[1];
            return `${lat.toFixed(4)}`;
          },
        },

        {
          title: "Longitude",
          render: (f) => {
            const lng = f.geometry.coordinates[0];
            return `${lng.toFixed(4)}`;
          },
        },
      ]}
    />
  );
}