<template>
  <div>
    <div>MyApp</div>
    <button v-on:click="showAlert">This is test</button>
  </div>
</template>
<script>
const appData = {
  count: 0
};

export default {
  data() {
    return appData;
  },
  methods: {
    showAlert: showAlert
  }
};

function showAlert() {
  console.log("test");
  alert("test");
}
</script>
<style>
</style>
