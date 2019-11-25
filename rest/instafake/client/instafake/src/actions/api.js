import axios from "axios";

const api = axios.default.create({
  baseURL: "http://localhost:5000/gw/"
});

export const user = (id) => api.get(`users/${id}`);