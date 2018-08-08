var app = new Vue({
    el: '#app',
    delimiters: ['${', '}'],
    data: {
        searchInput: '',
        searchType: '',
        searchOutput: {},
        tableHeaderInfo: {},
    },
    computed: {
        searchButtonValue: function () {
            if (this.searchType === 'Earnings' && this.searchInput !== '')
                return 'Get Earnings';
            else if (this.searchType === 'Dividends' && this.searchInput !== '')
                return 'Get Dividends';
            else
                return 'Disabled';
        },
        isDisabled: function () {
            return (this.searchType === '' || this.searchInput === '')
        }
    },
    methods: {
        sendData() {
            if (this.searchType === "Earnings")
                getEarnings(this);
            else if (this.searchType === "Dividends")
                getDividends(this);
        },
        addRowToDB(res) {
            if (res == undefined) return;
            if (this.searchType === "Earnings") {
                alert("Stock Earnings.");
                return
            } else  if (this.searchType === "Dividends") {
                alert("Stock Dividend.");
                return
            }
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

