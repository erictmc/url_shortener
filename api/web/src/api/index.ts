import axios from "axios";
import {
    convertObjToCamelCase,
    convertObjToSnakeCase
} from "../utils";


axios.interceptors.request.use((config) => {
    let { data } = config;
    try {
        data = convertObjToSnakeCase(data || {});
    } catch (e) {
        console.warn(e);
    }
    return {
        ...config,
        data,
    };
});

axios.interceptors.response.use((response) => {
    return {
        ...response,
        data: convertObjToCamelCase(response.data || {}),
    };
});
