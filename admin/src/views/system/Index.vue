<template>
  <el-card>
    <el-tabs v-model="activeName" @tab-click="handleClick">
      <el-tab-pane v-for="item in tabs" :label="item.label" :name="item.name">
      </el-tab-pane>
    </el-tabs>
    <component :is="tabs[activeName].component" />
  </el-card>
</template>

<script setup>
import { ref } from "vue";
import Currency from "./currency.vue";
const activeName = ref("currency");
let tabs = {
  currency: { label: "通用配置", name: "currency", component: Currency },
};
const handleClick = (tab, event) => {
  console.log(tab, event);
};
</script>
