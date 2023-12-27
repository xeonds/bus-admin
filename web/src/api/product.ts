import serviceAxios from "../axios";

const productAPI = {
  BASE_URL: '/products',
  getProducts: async () => serviceAxios.get(productAPI.BASE_URL).then(response => response.data),
  getProductById: async (id: number) => serviceAxios.get(`${productAPI.BASE_URL}/${id}`).then(response => response.data),
  createProduct: async (product: any) => serviceAxios.post(productAPI.BASE_URL, product).then(response => response.data),
  updateProduct: async (id: number, product: any) => serviceAxios.put(`${productAPI.BASE_URL}/${id}`, product).then(response => response.data),
  deleteProduct: async (id: number) => serviceAxios.delete(`${productAPI.BASE_URL}/${id}`).then(response => response.data),
};

export default productAPI;