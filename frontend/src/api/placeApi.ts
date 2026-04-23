import axios from "axios";
import { Feature } from "../interfaces/place";

const api = axios.create({
  // baseURL: "http://localhost:8080/api",
  baseURL: import.meta.env.VITE_API_URL + "/api",
  timeout: 5000,
  headers: {
    "Content-Type": "application/json",
  },
});


// GET all places
export const getFeatures = async (
  search?: string,
  category?: string
): Promise<Feature[]> => {
  const res = await api.get("/collections/places/items", {
    params: {
      search,
      category,
    },
  });

  return res.data?.features ?? [];
};

// CREATE 
export const createFeature = async (data: any) => {
  return api.post("/collections/places/items", data);
};

// DELETE 
export const deleteFeature = async (id: string) => {
  return api.delete(`/collections/places/items/${id}`);
};
//Update 
export const updateFeature = async (id: string, data: any) => {
  return api.patch(`/collections/places/items/${id}`, data);
};

export const getCategories = async () => {
  const res = await api.get("/categories");
  return res.data;
};

export const getPlacesWithin = async (polygon: any) => {
  const res = await api.post("/collections/places/within", {
    polygon,
  });

  return res.data?.features ?? [];
};