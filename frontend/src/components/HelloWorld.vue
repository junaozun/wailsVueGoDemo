<script setup>
import {reactive, ref} from 'vue'
import {Greet} from '../../wailsjs/go/main/App'
import {ConnectionList} from '../../wailsjs/go/main/App'
import {ConnectionCreate} from '../../wailsjs/go/main/App'

const data = reactive({
  name: "",
  resultText: "Please enter redis Addr:",
})

let code = ref()
let list = ref()
let createRes = ref()

function greet() {
  // Greet(data.name).then(result => {
  //   data.resultText = result
  // })
  ConnectionCreate({"addr":data.name}).then(res => {
    createRes.value = res
  })
  ConnectionList().then(res => {
    code.value = res.code
    list.value = res.data
  })
}

</script>

<template>
  <main>
    <div id="result" class="result">{{ data.resultText }}</div>
    <div id="input" class="input-box">
      <input id="name" v-model="data.name" autocomplete="off" class="input" type="text"/>
      <button class="btn" @click="greet">Greet</button>
      <div>code :{{code}}</div>
      <div>list :{{list}}</div>
      <div>CreateRes :{{createRes}}</div>
    </div>
  </main>
</template>

<style scoped>
.result {
  height: 20px;
  line-height: 20px;
  margin: 1.5rem auto;
}

.input-box .btn {
  width: 60px;
  height: 30px;
  line-height: 30px;
  border-radius: 3px;
  border: none;
  margin: 0 0 0 20px;
  padding: 0 8px;
  cursor: pointer;
}

.input-box .btn:hover {
  background-image: linear-gradient(to top, #cfd9df 0%, #e2ebf0 100%);
  color: #333333;
}

.input-box .input {
  border: none;
  border-radius: 3px;
  outline: none;
  height: 30px;
  line-height: 30px;
  padding: 0 10px;
  background-color: rgba(240, 240, 240, 1);
  -webkit-font-smoothing: antialiased;
}

.input-box .input:hover {
  border: none;
  background-color: rgba(255, 255, 255, 1);
}

.input-box .input:focus {
  border: none;
  background-color: rgba(255, 255, 255, 1);
}
</style>
