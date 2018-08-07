var app = new Vue({
    el: '#app',
    delimiters: ['${', '}'],
    data: {
        searchInput: '',
        searchType: '',
        searchOutput: {},
        tableHeaderInfo: {}
    },
    methods: {
        sendData() {
            if (this.searchType === "Books")
                searchBooks(this)
            else if (this.searchType === "Stocks")
                getEarnings(this)
        },
        addRowToDB(res) {
            if (res == undefined) return;
            if (this.searchType === "Stocks") {
                alert("Nothing to do with stocks.")
                return
            }
            this.$http.post('/books/add', res).then((response) => {
                alert("added the record");
            });
        }
    }
});

Vue.component('datacell', {
    delimiters: ['${', '}'],
    props: {
        widthPercent: String,
        value: String,
        cssClassName: String
    },
    template: '<td :width="widthPercent" :class="cssClassName"> ${value}</td>'
});

Vue.component('headercell', {
    delimiters: ['${', '}'],
    props: {
        widthPercent: String,
        value: String,
        cssClassName: String
    },
    template: '<th :width="widthPercent" :class="cssClassName"> ${value}</th>'
});

