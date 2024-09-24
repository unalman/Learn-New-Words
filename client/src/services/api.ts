import axios from "axios";
axios.defaults.headers.post['Content-Type'] ='application/json';
axios.defaults.headers.post['Access-Control-Allow-Origin'] = '*';

const instance = axios.create({
  baseURL: import.meta.env.VITE_API_ENDPOINT,
});
export default instance;
