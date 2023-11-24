// composables/useSize.js
export default function useSize() {
    const getSize = (bytes) => {
        // b 格式化成 m
       return (bytes/1024/1024).toFixed(2)
    }
    return { getSize };
}
