import axios from "axios";
import { Feature } from "../interfaces/place";

// =====================
// axios instance
// =====================
const api = axios.create({
  baseURL: "http://localhost:8080/api",
  timeout: 5000,
  headers: {
    "Content-Type": "application/json",
  },
});

// =====================
// Place APIs
// =====================

// GET all places
export const getFeatures = async (): Promise<Feature[]> => {
  const res = await api.get("/collections/places/items");
  return res.data.features;
};

// // GET by id
// export const getPlaceById = async (id: string): Promise<Place> => {
//   const res = await api.get(`/places/${id}`);
//   return res.data.data;
// };

// CREATE place
export const createFeature = async (data: any) => {
  return api.post("/collections/places/items", data);
};

// DELETE place
export const deleteFeature = async (id: string) => {
  return api.delete(`/collections/places/items/${id}`);
};