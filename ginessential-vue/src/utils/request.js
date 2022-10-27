import storageService from "@/service/storageService";
import axios from "axios";

const service = axios.create({
    baseURL: 'http://localhost:1016/api/',
    timeout: 1000 * 5,
    headers: { Authorization: `Bearer ${storageService.get(storageService.USER_TOKEN)}` }
});

service.interceptors.request.use(function (config) {
    // Do something before request is sent
    Object.assign(config.headers, { Authorization: `Bearer ${storageService.get(storageService.USER_TOKEN)}` })
    return config;
}, function (error) {
    // Do something with request error
    return Promise.reject(error);
});

export default service;