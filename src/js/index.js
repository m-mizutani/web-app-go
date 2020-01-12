import '../css/main.scss'

import _ from 'lodash';
import "babel-polyfill";
import Vue from "vue";

import App from "./app.vue";

new Vue({
    el: "#app",
    render: h => h(App)
});
