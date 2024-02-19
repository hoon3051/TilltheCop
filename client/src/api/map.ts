import { LocationParams } from "../model/location";
import ApiManager from ".";

export const getMap = async (location: LocationParams): Promise<{ map_url: string }> => {
  const response = await ApiManager.post("/map", location);
  return response.data;
};
