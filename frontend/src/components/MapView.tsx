import { useEffect, useRef } from "react";
import maplibregl from "maplibre-gl";
import "maplibre-gl/dist/maplibre-gl.css";

export default function MapView({ features }: any) {
  const mapRef = useRef<maplibregl.Map | null>(null);
  const markersRef = useRef<maplibregl.Marker[]>([]);

  useEffect(() => {
    const map = new maplibregl.Map({
      container: "map",
      style: "https://tiles.openfreemap.org/styles/liberty",
      center: [100.5231, 13.7367],
      zoom: 10,
    });

    mapRef.current = map;

    map.on("load", () => {
      // 🧹 เคลียร์ marker เก่า (สำคัญเวลา state เปลี่ยน)
      markersRef.current.forEach((m) => m.remove());
      markersRef.current = [];

      // 📍 สร้าง marker
      features.forEach((f: any) => {
        const marker = new maplibregl.Marker()
          .setLngLat(f.geometry.coordinates)
          .setPopup(
            new maplibregl.Popup().setText(f.properties.name)
          )
          .addTo(map);

        markersRef.current.push(marker);
      });
    });

    return () => {
      markersRef.current.forEach((m) => m.remove());
      map.remove();
    };
  }, [features]);

  return <div id="map" style={{ height: "100%", width: "100%" }} />;
}