var myHtml = `
<h1>欢迎使用GoFast示例程序！</h1>
<button @click="count++">
    You clicked me {{ count }} times.
</button>
`
export default {
    template: myHtml,
    setup() {
        const count = Vue.ref(0)
        return {count}
    },
}