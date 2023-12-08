<template>
  <div>
    <CardTitleCard
      title="更多分组"
      v-model="actived"
      :list="data"
    />
    <!-- <CardTitleCard title="More Pages" class="mb-5 h-80" /> -->
  </div>
</template>

<script setup>
import { useGroupStore } from "~/stores/init.js";
const store = useGroupStore();
const data = computed(() => store.getGroup);
const actived = computed({
  get: () => store.actived,
  set: (val) => {
    store.setActived(val);
  },
});

store.fetchGroup();
</script>
