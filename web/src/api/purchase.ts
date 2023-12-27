import serviceAxios from "../axios";

const purchaseAPI = {
  BASE_URL: '/purchases',
  getPurchases: async () => serviceAxios.get(purchaseAPI.BASE_URL).then(response => response.data),
  getPurchaseById: async (id: number) => serviceAxios.get(`${purchaseAPI.BASE_URL}/${id}`).then(response => response.data),
  createPurchase: async (purchase: any) => serviceAxios.post(purchaseAPI.BASE_URL, purchase).then(response => response.data),
  updatePurchase: async (id: number, purchase: any) => serviceAxios.put(`${purchaseAPI.BASE_URL}/${id}`, purchase).then(response => response.data),
  deletePurchase: async (id: number) => serviceAxios.delete(`${purchaseAPI.BASE_URL}/${id}`).then(response => response.data),
};

export default purchaseAPI;