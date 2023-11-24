<template>
  <div>
    <el-form :model="form" label-width="150px">
      <el-form-item label="Activity name">
        <el-input v-model="form.name" />
      </el-form-item>
    </el-form>
  </div>
</template>

<script setup>
import { ref } from "vue";
let form = ref({
  name: "",
});
</script>
