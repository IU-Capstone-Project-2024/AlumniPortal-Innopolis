import { createApp } from 'vue'
import { createRouter, createWebHistory } from 'vue-router'

import HelloWorld from './views/HomeView.vue'

const router = createRouter({
  history: createWebHistory(),
  routes: [
    {
      path: '/',
      name: 'Home',
      component: HelloWorld,
    },
  ]
})

import './style.css'
import App from './App.vue'

const app = createApp(App)

app.use(router)
app.mount('#app')
