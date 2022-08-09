import { createRouter, createWebHistory } from 'vue-router'
import Login from '../src/components/login.vue'
import Product from '../src/components/product.vue'
import SignUp from '../src/components/signUp.vue'
import ShoppingCart from '../src/components/shoppingCart.vue'
import Order from '../src/components/order.vue'
import Tmp from '../src/components/tmp.vue'
import Manager from '../src/components/manager.vue'
let history = createWebHistory()
let routes = [
  {
    path: '/',
    name: 'Product',
    component: Product
  },
  {
    path: '/Login',
    name: 'Login',
    component: Login
  },
  {
    path: '/SignUp',
    name: 'SignUp',
    component: SignUp
  },
  {
    path: '/ShoppingCart',
    name: 'ShoppingCart',
    component: ShoppingCart
  },
  {
    path: '/Manager',
    name: 'Manager',
    component: Manager
  },
  {
    path: '/Order',
    name: 'Order',
    component: Order
  },
  {
    path: '/Tmp',
    name: 'Tmp',
    component: Tmp
  },
]

export default createRouter({ history, routes })