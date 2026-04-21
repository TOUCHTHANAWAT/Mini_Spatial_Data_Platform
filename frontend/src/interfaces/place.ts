export interface Feature {
  type: "Feature";
  id: string;
  geometry: {
    type: "Point";
    coordinates: [number, number];
  };
  properties: {
    name: string;
  };
}

export interface FeatureCollection {
  type: "FeatureCollection";
  features: Feature[];
}