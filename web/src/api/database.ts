import serviceAxios from '../axios';

const databaseAPI = {
  BASE_URL: '/database',
  backup: async () => serviceAxios.post(`${databaseAPI.BASE_URL}/backup`).then(response => response.data),
  restore: async () => serviceAxios.post(`${databaseAPI.BASE_URL}/restore`).then(response => response.data),
};

export default databaseAPI;
