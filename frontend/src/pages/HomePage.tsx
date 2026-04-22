import { useEffect, useState } from "react";
import { getFeatures, getCategories } from "../api/placeApi";
import MapView from "../components/MapView";
import { Input, List, Drawer, Button, message, Modal, Segmented } from "antd";
import * as maplibregl from "maplibre-gl";
import { Feature, Category } from "../interfaces/place";
import { deleteFeature } from "../api/placeApi";
import EditPlaceModal from "../components/EditPlaceModal";
import { SearchOutlined } from '@ant-design/icons';
import { getPlacesWithin } from "../api/placeApi";

export default function HomePage() {
  const [features, setFeatures] = useState<Feature[]>([]);
  const [search, setSearch] = useState("");
  const [map, setMap] = useState<maplibregl.Map | null>(null);
  const [open, setOpen] = useState(false);
  const [categories, setCategories] = useState<Category[]>([]);
  const [category, setCategory] = useState<string>("");
  const [editingFeature, setEditingFeature] = useState<Feature | null>(null);
  const [drawResults, setDrawResults] = useState<Feature[]>([]);

  useEffect(() => {
    const timeout = setTimeout(async () => {
      try {
        const data = await getFeatures(search, category);
        setFeatures(data ?? []);
      } catch (err) {
        console.error(err);

        message.error("Location data failed to load.");
      }
    }, 300);

    return () => clearTimeout(timeout);
  }, [search, category]);

  useEffect(() => {
    const fetchCategories = async () => {
      try {
        const data = await getCategories();
        setCategories(data ?? []);
      } catch (err) {
        console.error(err);

        message.error("Category loading failed.");
      }
    };

    fetchCategories();
  }, []);

  const flyToFeature = (feature: Feature) => {
    if (!map) return;

    map.flyTo({
      center: feature.geometry.coordinates,
      zoom: 15,
      speed: 1.2,
    });
  };

  const handleEdit = (feature: Feature) => {
    setEditingFeature(feature);
  };

  const handleDelete = (id: string) => {
    Modal.confirm({
      title: "Delete this location?",
      okType: "danger",

      onOk: async () => {
        try {
          await deleteFeature(id);
          message.success("Deleted successfully");

          const data = await getFeatures(search, category);
          setFeatures(data ?? []);
        } catch {
          message.error("Delete failed");
        }
      },
    });
  };

  const handleDrawSearch = async (polygon: any) => {
    try {
      if (!polygon) {
        setDrawResults([]);
        const data = await getFeatures(search, category);
        setFeatures(data ?? []);
        return;
      }

      const data = await getPlacesWithin(polygon);
      console.log(data)
      setDrawResults(data ?? []);
      setFeatures(data);
    } catch (err) {
      message.error("Search area failed");
    }
  };

  return (
    <div style={{ height: "100vh", position: "relative" }}>
      <div
        style={{
          position: "absolute",
          top: 20,
          right: 50,
          zIndex: 1000,
          display: "flex",
          alignItems: "center",
          gap: 8,
          background: "#fff",
          padding: "6px 8px",
          borderRadius: 999,
          boxShadow: "0 4px 12px rgba(0,0,0,0.15)",
        }}
      >
        <Segmented
          size="small"
          value={category || "all"}
          onChange={(val) =>
            setCategory(val === "all" ? "" : String(val))
          }
          options={[
            { label: "All", value: "all" },
            ...categories.map((c) => ({
              label: c.label,
              value: c.value,
            })),
          ]}
        />
        <Button
          shape="circle"
          icon={<SearchOutlined />}
          onClick={() => setOpen(true)}
        />
      </div>

      <Drawer
        title="Location"
        placement="right"
        open={open}
        onClose={() => setOpen(false)}
        width={320}
        mask={false}
      >
        <Input
          placeholder="Search Location..."
          allowClear
          value={search}
          onChange={(e) => setSearch(e.target.value)}
        />

        <List
          style={{ marginTop: 10 }}
          dataSource={features ?? []}
          renderItem={(item) => (
            <List.Item
              style={{ cursor: "pointer" }}
              onClick={() => {
                flyToFeature(item);
                setOpen(false);
              }}
            >
              {item.properties.name}
            </List.Item>
          )}
        />
      </Drawer>
      <EditPlaceModal
        open={!!editingFeature}
        feature={editingFeature}
        categories={categories}
        onClose={() => setEditingFeature(null)}
        onSuccess={async () => {
          const data = await getFeatures(search, category);
          setFeatures(data ?? []);
        }}
      />

      {drawResults.length > 0 && (
        <div
          style={{
            position: "absolute",
            top: 20,
            left: 20,
            zIndex: 1000,
            width: 260,
            maxHeight: 300,
            overflowY: "auto",
            background: "#fff",
            borderRadius: 12,
            padding: 10,
            boxShadow: "0 4px 12px rgba(0,0,0,0.15)",
          }}
        >
          <div style={{ fontWeight: 600, marginBottom: 8 }}>
            Location in selected area ({drawResults.length})
          </div>

          <List
            size="small"
            dataSource={drawResults}
            renderItem={(item) => (
              <List.Item
                style={{ cursor: "pointer" }}
                onClick={() => flyToFeature(item)}
              >
                {item.properties.name}
              </List.Item>
            )}
          />
        </div>
      )}

      <div style={{ height: "100%", width: "100%" }}>
        <MapView features={features} categories={categories} onMapReady={setMap} onEdit={handleEdit} onDelete={handleDelete} onDrawSearch={handleDrawSearch} />
      </div>
    </div>
  );
}