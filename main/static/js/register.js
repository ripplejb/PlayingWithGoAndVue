var registerApp = new Vue({
    el: '#registerApp',
    delimiters: ['${', '}'],
    data: {
        username: '',
        password1: '',
        password2: '',
        error: ''
    },
    template:
        '<form style="display: flex; align-items: center; justify-content: center; flex-flow: column; height: 350px; width: 100%;">\n' +
        '    <div style="display: flex; align-items: center; justify-content: center; flex-flow: column; height: 200px; width: auto;">\n' +
        '        <div style="align-self: flex-start; width:100%; display: flex; justify-content: center;">\n' +
        '            <h4>Register New User</h4>\n' +
        '        </div>\n' +
        '        <div style="align-self: flex-start; width:auto; display: flex; align-items: baseline; ">\n' +
        '            <label for="user-id" style="width: 100px;">Email: </label><input type="email" ref="email" name="user-id" v-model="username" required>\n' +
        '        </div>\n' +
        '        <div style="align-self: flex-start; width: auto; display: flex; align-items: baseline; margin-top: 10px;">\n' +
        '            <label for="password" style="width: 100px;">Password: </label><input type="password" ref="password1" name="password" v-model="password1" required>\n' +
        '        </div>\n' +
        '        <div style="align-self: flex-start; width: auto; display: flex; align-items: baseline; margin-top: 10px;">\n' +
        '            <label for="password" style="width: 100px;">Password: </label><input type="password" ref="password2" name="password" v-model="password2" required>\n' +
        '        </div>\n' +
        '        <div style="align-self: flex-start; width: 100%; display: flex; justify-content: right; margin-top: 10px;">\n' +
        '            <input type="button" name="register" value="Register" v-on:click="register">\n' +
        '        </div>\n' +
        '    </div>\n' +
        '    <div style="color: red; margin-top: 30px; height: 50px;">${error}</div>\n' +
        '</form>\n',
    methods: {
        register: function (event) {
            var re = /^(([^<>()[\]\\.,;:\s@\"]+(\.[^<>()[\]\\.,;:\s@\"]+)*)|(\".+\"))@((\[[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\])|(([a-zA-Z\-0-9]+\.)+[a-zA-Z]{2,}))$/;
            if (!re.test(this.username)) {
                this.error = 'invalid email id.';
                this.$refs.email.focus();
                return;
            }

            if (this.password1 !== this.password2) {
                this.error = 'password do not match';
                this.$refs.password1.focus();
                return;
            }
            this.$http.post('/register', {"username": this.username, "password": this.password2})
                .then(
                    response => {
                        window.location.href = '/login-ui';
                    },
                    response => {
                        this.error = 'error creating new user'
                    });

        },

    }
});

