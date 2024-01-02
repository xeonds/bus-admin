import { createStore } from 'vuex'
import auth from './modules/auth'

// store token when it is set
const localStoragePlugin = (store) => {
  store.subscribe((mutation, { auth }) => {
    if (mutation.type === 'auth/SET_TOKEN') {
      localStorage.setItem('token', auth.token)
    }
  })
}

const store = createStore({
  plugins: [localStoragePlugin],
  modules: {
    auth,
  },
})

export default store
