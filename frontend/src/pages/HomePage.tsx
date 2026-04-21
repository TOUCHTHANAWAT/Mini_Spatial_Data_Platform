import { useEffect, useState } from "react";
import { getFeatures } from "../api/placeApi";
import MapView from "../components/MapView";

export default function HomePage() {
  const [features, setFeatures] = useState<any[]>([]);

  useEffect(() => {
    const fetch = async () => {
      try {
        const data = await getFeatures();
        setFeatures(data);
        // console.log(data)
      } catch (err) {
        console.error(err);
      }
    };

    fetch();
  }, []);

  return (
    <div style={{ height: "85vh", display: "flex", flexDirection: "column" }}>
      <h2 style={{ margin: 0, padding: 10 }}>Map Overview</h2>

      <div style={{ flex: 1 }}>
        <MapView features={features} />
      </div>
    </div>
  );
}