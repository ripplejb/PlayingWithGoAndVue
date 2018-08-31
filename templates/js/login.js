var loginApp = new Vue({
    el: '#loginApp',
    delimiters: ['${', '}'],
    data: {
        email: '',
        password: '',
        error: ''
    },
    methods: {
        registerUser: function (event) {
            window.location.href = '/register-ui'
        },

        login: function (event) {
            this.$http.post('/login', {"username": this.username, "password": this.password2}).then(
                response => {
                    this.error = response.toString()
                },
                response => {
                    this.error = response.body;
                });
        }
    },
    template: '<form style="display: flex; align-items: center; justify-content: center; flex-flow: column; height: 200px; width: 100%;">\n' +
        '    <div style="display: flex; align-items: center; justify-content: center; flex-flow: column; height: 200px; width: auto;">\n' +
        '        <div style="align-self: flex-start; width:100%; display: flex; justify-content: center;">\n' +
        '            <h4>Please Sign In</h4>\n' +
        '        </div>\n' +
        '        <div style="align-self: flex-start; width:auto; display: flex; align-items: baseline; margin-bottom: 5px;">\n' +
        '            <label for="user-id" style="width: 100px;">User Name: </label><input type="email" name="user-id" required>\n' +
        '        </div>\n' +
        '        <div style="align-self: flex-start; width: auto; display: flex; align-items: baseline; margin-top: 5px;">\n' +
        '            <label for="password" style="width: 100px;">Password: </label><input type="password" name="password" required>\n' +
        '        </div>\n' +
        '        <div style="align-self: flex-start; width: 100%; display: flex; justify-content: space-evenly; margin-top: 10px;">\n' +
        '            <input type="button" name="register" value="Register" v-on:click="registerUser">\n' +
        '            <input type="button" name="sign-in" value="Sign In" v-on:click="login">\n' +
        '        </div>\n' +
        '    </div>\n' +
        '    <div style="color: red; margin-top: 10px;">${error}</div>\n' +
        '</form>\n'
});