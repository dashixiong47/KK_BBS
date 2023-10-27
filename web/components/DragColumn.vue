<script setup>
import Sortable from "sortablejs";

const myList = ref(null);
let props = defineProps({
  items: {
    type: Array,
    default: () => [],
  },
});
const items = ref(props.items);

onMounted(() => {
  new Sortable(myList.value, {
    animation: 150,
    ghostClass: "sortable-ghost",
    onUpdate: function (evt) {
      const oldIndex = evt.oldIndex;
      const newIndex = evt.newIndex;
      const movedItem = items.value.splice(oldIndex, 1)[0];
      items.value.splice(newIndex, 0, movedItem);
    },
  });
});
</script>

<template>
  <div>
    <ul ref="myList" class="glass rounded-xl">
      <li
        v-for="(item, index) in items"
        :key="item"
        class="flex p-4"
        :class="[!Object.is(index,0) || !Object.is(index,items.length-1)?'border-b':'',]"
      >
        <div class="min-w-[50px]">{{ index+1 }}</div>
        <div class="w-24 ">{{ item.name }}</div>
        <div class="w-24 ">{{ item.type }}</div>
        <div class="w-24 ">{{ item.size }}</div>
      </li>
    </ul>
  </div>
  
</template>
