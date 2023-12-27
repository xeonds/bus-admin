import serviceAxios from "../axios";

const orderAPI = {
  BASE_URL: '/orders',
  getOrders: async () => serviceAxios.get(orderAPI.BASE_URL).then(response => response.data),
  getOrdersToday: async () => serviceAxios.get(`${orderAPI.BASE_URL}/today`).then(response => response.data),
  getOrderById: async (id: number) => serviceAxios.get(`${orderAPI.BASE_URL}/${id}`).then(response => response.data),
  createOrder: async (order: any) => serviceAxios.post(orderAPI.BASE_URL, order).then(response => response.data),
  updateOrder: async (id: number, order: any) => serviceAxios.put(`${orderAPI.BASE_URL}/${id}`, order).then(response => response.data),
  deleteOrder: async (id: number) => serviceAxios.delete(`${orderAPI.BASE_URL}/${id}`).then(response => response.data),
};

export default orderAPI;