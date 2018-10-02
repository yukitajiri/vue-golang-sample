// The Vue build version to load with the `import` command
// (runtime-only or standalone) has been set in webpack.base.conf with an alias.
import Vue from 'vue'
import App from './App'
import router from './router'
import firebase from 'firebase'

Vue.config.productionTip = false

/* eslint-disable no-new */
new Vue({
  el: '#app',
  router,
  components: { App },
  template: '<App/>'
})


Vue.config.productionTip = false

const config = {
  apiKey: "AIzaSyCSQtWzLbWG5DwV7W50Xjf9m1g1FwyAjvg",
  authDomain: "golang-app-183309.firebaseapp.com",
};
firebase.initializeApp(config)
