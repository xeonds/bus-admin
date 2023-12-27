import serviceAxios from '../axios';

const authAPI = {
  BASE_URL: '/auth',
  login: async (username: string, password: string) => serviceAxios.post(`${authAPI.BASE_URL}/login`, { username, password }).then(response => response.data),
};

export default authAPI;
