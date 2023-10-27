<template>
   <div class="w-[800px] h-96">
      <DragColumn :items="modelValue"></DragColumn>
      <Upload
        accept="*"
        multiple
        @uploadSuccess="uploadSuccess"
      >
        <KButton>上传</KButton>
      </Upload>
    </div>
</template>

<script setup>
let uploadGattachmentList = ref([]);
const props=defineProps({
  modelValue:{
    type:Array,
    default:()=>[]
  }
})
const emit=defineEmits(["update:modelValue"])
const modelValue=computed({
  get:()=>{
    if(props.modelValue){
      return props.modelValue
    }
    return uploadGattachmentList.value
  },
  set:(value)=>{
    console.log(value);
    if(props.modelValue){
      emit("update:modelValue",value)
      return
    }
    uploadGattachmentList.value=value
  }
})
const uploadSuccess = (data) => {
  modelValue.value.push(data);
};
</script>