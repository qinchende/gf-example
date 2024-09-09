var loginHtml = `
<div>
  <template v-if="showPanel">
  <table class="logTable">
  <thead><tr><td>操作</td><td>时间</td><td>IP地址</td><td>媒介</td></tr></thead>
  <tbody></tbody>
  </table>
  <hr />
  <p>RetData: {{ retData }}</p>
  </template>
</div>
`
export default {
    template: loginHtml,
    setup() {
    },
    data() {
        return {
            showPanel: true,
            retData: 'NA'
        }
    },
    methods: {
        onLogin() {
            axios.get(GFDemo.Host + '/request_test_data')
                .then(resp => (this.retData = resp.data))
                .catch(err => this.retData = err)
        }
    },
}