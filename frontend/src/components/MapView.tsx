import { useEffect, useRef } from "react";
import maplibregl from "maplibre-gl";
import "maplibre-gl/dist/maplibre-gl.css";
import { Feature, Category } from "../interfaces/place";
import "./MapView.css"
import MapboxDraw from "maplibre-gl-draw";
import "@mapbox/mapbox-gl-draw/dist/mapbox-gl-draw.css";

type Props = {
  features: Feature[];
  categories: Category[];
  onMapReady: (map: maplibregl.Map) => void;
  onEdit: (feature: Feature) => void;
  onDelete: (id: string) => void;
  onDrawSearch: (polygon: any) => void;
};

export default function MapView({ features, categories, onMapReady, onEdit, onDelete, onDrawSearch }: Props) {
  const mapRef = useRef<maplibregl.Map | null>(null);
  const markersRef = useRef<maplibregl.Marker[]>([]);

  useEffect(() => {
    const map = new maplibregl.Map({
      container: "map",
      style: "https://tiles.openfreemap.org/styles/liberty",
      // style: "https://demotiles.maplibre.org/style.json",
      center: [100.5231, 13.7367], // Bangkok
      zoom: 10,
    });

    mapRef.current = map;

    onMapReady(map);

    const draw = new MapboxDraw({
      displayControlsDefault: false,
      controls: {
        polygon: true,
        trash: true,
      },
    });

    map.addControl(draw as unknown as maplibregl.IControl);

    map.on("draw.create", (e: any) => {
      const polygon = e.features[0].geometry;
      console.log("Polygon:", polygon);
      onDrawSearch(polygon);
    });

    map.on("draw.delete", () => {
      onDrawSearch(null); 
    });

    return () => {
      map.remove();
    };
  }, []);
  
  useEffect(() => {
    const map = mapRef.current;
    if (!map) return;

    // remove old markers
    markersRef.current.forEach((m) => m.remove());
    markersRef.current = [];

    // add new markers
    features.forEach((f) => {
      const category = categories.find(
        (c) => c.value === f.properties.category
      );
      // 1. สร้าง DOM
      const popupNode = document.createElement("div");
      popupNode.className = "popup-wrapper";

      popupNode.innerHTML = `
    <div class="card">
      <div class="header">
        <h3>${f.properties.name}</h3>
                <span 
          class="badge"
          style="
            background: ${category?.color || "#e5e7eb"};
            color: white;
          "
        >
          ${category?.label || f.properties.category}
        </span>
      </div>

      <p class="coords">
        ${f.geometry.coordinates[1]} N, ${f.geometry.coordinates[0]} E
      </p>

      <div class="actions">
        <button class="edit-btn">Edit Data</button>
        <button class="delete-btn">Delete</button>
      </div>
    </div>
  `;

      // 2. bind event (สำคัญ!)
      popupNode.querySelector(".edit-btn")?.addEventListener("click", () => {
        onEdit(f); //เรียก HomePage
      });

      popupNode.querySelector(".delete-btn")?.addEventListener("click", () => {
        onDelete(f.id); //เรียก HomePage
      });

      // 3. สร้าง popup
      const popup = new maplibregl.Popup({
        offset: 25,
        className: "my-popup",
        closeButton: false,
      }).setDOMContent(popupNode);

      // 4. สร้าง marker
      const marker = new maplibregl.Marker()
        .setLngLat(f.geometry.coordinates)
        .setPopup(popup)
        .addTo(map);

      markersRef.current.push(marker);
    });
  }, [features, categories]);



  return <div id="map" style={{ height: "100%", width: "100%" }} />;
}