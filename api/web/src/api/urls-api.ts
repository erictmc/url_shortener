import {UrlForm, UrlResult} from "../models/urls";
import axios from "axios";

export const createUrl = async (form: UrlForm): Promise<UrlResult> => {
    const { data } = await axios.post("/url/new", form)
    return data
}