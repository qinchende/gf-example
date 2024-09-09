var loginHtml = `
<div>
  <template v-if="showPanel">
  <label for="account">账号：</label><input id="account" placeholder="请输入手机号">
  <br />
  <label for="passwd">密码：</label><input id="passwd" placeholder="6位以上字符">
  <br />
  <button v-on:click="onLogin" style="width: 120px; margin: 5px 0;">登&nbsp;录</button>
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
            axios.get(GFDemo.Host + '/request_url')
                .then(resp => (this.retData = resp.data))
                .catch(err => this.retData = err)
        }
    },
}