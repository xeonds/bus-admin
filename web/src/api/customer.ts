import serviceAxios from '../axios';

const customerAPI = {
  BASE_URL: '/customers',
  getCustomers: async () => serviceAxios.get(customerAPI.BASE_URL).then(response => response.data),
  getCustomerById: async (id: number) => serviceAxios.get(`${customerAPI.BASE_URL}/${id}`).then(response => response.data),
  createCustomer: async (customer: any) => serviceAxios.post(customerAPI.BASE_URL, customer).then(response => response.data),
  updateCustomer: async (id: number, customer: any) => serviceAxios.put(`${customerAPI.BASE_URL}/${id}`, customer).then(response => response.data),
  deleteCustomer: async (id: number) => serviceAxios.delete(`${customerAPI.BASE_URL}/${id}`).then(response => response.data)
};

export default customerAPI;
