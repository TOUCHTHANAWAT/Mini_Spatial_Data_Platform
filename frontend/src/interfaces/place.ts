export interface Feature {
  type: "Feature";
  id: string;
  geometry: {
    type: "Point";
    coordinates: [number, number];
  };
  properties: {
    name: string;
    category: string;
  };
}

export interface FeatureCollection {
  type: "FeatureCollection";
  features: Feature[];
}

export interface Category {
  value: string;
  label: string;
  color: string;
}