import { useEffect, useState } from "react";
import { Table, Button, message, Modal, Tag } from "antd";
import { getFeatures, deleteFeature } from "../api/placeApi";
import { Feature } from "../interfaces/place";
import { getCategories } from "../api/placeApi";
import { Category } from "../interfaces/place";
import "./PlacePage.css"
import CreatePlaceModal from "../components/CreatePlace";
import EditPlaceModal from "../components/EditPlaceModal";

export default function PlacePage() {
  const [features, setFeatures] = useState<Feature[]>([]);
  const [categories, setCategories] = useState<Category[]>([]);
  const [isCreateOpen, setIsCreateOpen] = useState(false);
  const [editingFeature, setEditingFeature] = useState<Feature | null>(null);

  const fetchData = async () => {
    try {
      const [featuresData, categoriesData] = await Promise.all([
        getFeatures(),
        getCategories(),
      ]);
      console.log(featuresData)
      setFeatures(featuresData);
      setCategories(categoriesData);
    } catch (err) {
      message.error("Data loading failed.");
      console.error(err);
    }
  };

  useEffect(() => {
    fetchData();
  }, []);


  const handleDelete = (id: string) => {
    Modal.confirm({
      title: "Delete this place?",
      content: "Are you sure you want to delete this item?",
      okText: "Delete",
      okType: "danger",
      cancelText: "Cancel",

      onOk: async () => {
        try {
          await deleteFeature(id);
          message.success("Deleted successfully");
          fetchData(); // reload data
        } catch (err) {
          message.error("Delete failed");
        }
      },
    });
  };

  const openEdit = (record: Feature) => {
    setEditingFeature(record);
  };

  const formatLatLng = (value: number, type: "lat" | "lng") => {
    const abs = Math.abs(value).toFixed(4);

    if (type === "lat") {
      return `${abs}° ${value >= 0 ? "N" : "S"}`;
    }

    return `${abs}° ${value >= 0 ? "E" : "W"}`;
  };

  return (
    <div>
      <div style={{ display: "flex", justifyContent: "space-between", marginBottom: 12 }}>
        <div>
          <h2>Spatial Data Registry</h2>
          <p>Manage and refine geographic nodes.</p>
        </div>

        <Button type="primary" onClick={() => setIsCreateOpen(true)}>
          Create Location
        </Button>
      </div>
      <Table
        className="custom-table"
        dataSource={features}
        rowKey="id"
        columns={[
          { title: "LOCATION NAME", render: (f: Feature) => f.properties.name },

          {
            title: "CATEGORY",
            align: "center",
            render: (f: Feature) => {
              const category = categories.find(
                (c) => c.value === f.properties.category
              );

              if (!category) return "-";

              return (
                <Tag color={category.color} style={{
                  borderRadius: 999,
                  padding: "2px 10px",
                  fontSize: 12,
                  fontWeight: 500,
                  display: "inline-flex",
                  alignItems: "center",
                  gap: 6,
                }}>
                  {category.label}
                </Tag>
              );
            },
          },

          {
            title: "LATITUDE",
            align: "center",
            render: (f: Feature) => formatLatLng(f.geometry?.coordinates?.[1], "lat"),
          },

          {
            title: "LONGITUDE",
            align: "center",
            render: (f: Feature) => formatLatLng(f.geometry?.coordinates?.[0], "lng"),
          },
          {
            title: "ACTION",
            align: "center",
            render: (_, record) => (
              <>
                <Button onClick={() => openEdit(record)} style={{ marginRight: 8 }}>
                  Edit
                </Button>

                <Button danger onClick={() => handleDelete(record.id)}>
                  Delete
                </Button>
              </>
            ),
          },
        ]}
      />
      <EditPlaceModal
        open={!!editingFeature}
        feature={editingFeature}
        categories={categories}
        onClose={() => setEditingFeature(null)}
        onSuccess={fetchData}
      />
      <CreatePlaceModal
        open={isCreateOpen}
        onClose={() => setIsCreateOpen(false)}
        onSuccess={fetchData}
      />
    </div>
  );
}