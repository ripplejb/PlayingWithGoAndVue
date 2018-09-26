Vue.component('datagrid', {
    delimiters: ['${', '}'],
    props: {
      tableHeaderInfo: Object,
      searchOutput: Object
    },
    template: '<div class="myTableContainer">\n' +
        '                <div class="myTable">\n' +
        '                    <table class="myheadertable">\n' +
        '                        <thead>\n' +
        '                        <tr class="mydataheaderrow">\n' +
        '                            <th is="headercell" v-for="hc in tableHeaderInfo" v-bind="hc"></th>\n' +
        '                        </tr>\n' +
        '                        </thead>\n' +
        '                    </table>\n' +
        '                    <table class="mydatatable">\n' +
        '                        <tbody id="search-result">\n' +
        '                        <tr class="mydatarow" v-for="res in searchOutput.Results" v-on:dblclick="addRowToDB(res)">\n' +
        '                            <td is="datacell" v-for="dc in res.celldef" v-bind="dc"></td>\n' +
        '                        </tr>\n' +
        '                        </tbody>\n' +
        '                    </table>\n' +
        '                </div>\n' +
        '            </div>'
});

Vue.component('dataform', {
    delimiters: ['${', '}'],
    data: function() {
        return {
            searchInput: '',
            searchType: '',
            tableHeaderInfo: {},
            searchOutput: {}
        }
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
            if (res === undefined) {
                return;
            }
            if (this.searchType === "Earnings") {
                alert("Stock Earnings.");
            } else  if (this.searchType === "Dividends") {
                alert("Stock Dividend.");
            }
        }
    },
    template:
        '            <form id="form-group" onsubmit="return false">\n' +
        '                <div class="form-grid-3area">\n' +
        '                    <div class="form-grid-search-row">\n' +
        '                        <input type="text" name="search" v-model="searchInput">\n' +
        '                        <input type="submit" :value="searchButtonValue" v-on:click="sendData" :disabled="isDisabled">\n' +
        '                    </div>\n' +
        '                    <div class="form-grid-option1">\n' +
        '                        <label class="radiocontainer" for="earnings">Stock Earnings<input type="radio"\n' +
        '                                                                                          name="searchOption"\n' +
        '                                                                                          id="earnings"\n' +
        '                                                                                          value="Earnings"\n' +
        '                                                                                          v-model="searchType">\n' +
        '                            <span class="radiocheckmark"></span>\n' +
        '                        </label>\n' +
        '                    </div>\n' +
        '                    <div class="form-grid-option2">\n' +
        '                        <label class="radiocontainer" for="dividends">Stock Dividends<input type="radio"\n' +
        '                                                                                            name="searchOption"\n' +
        '                                                                                            id="dividends"\n' +
        '                                                                                            value="Dividends"\n' +
        '                                                                                            v-model="searchType">\n' +
        '                            <span class="radiocheckmark"></span>\n' +
        '                        </label>\n' +
        '                    </div>\n' +
        '                </div>\n' +
        '            </form>\n'
});

var app = new Vue({
    el: '#app',
    delimiters: ['${', '}'],
    data: {
        searchOutput: {},
        tableHeaderInfo: {},
    },
    methods: {
        dataReceived: function (tableHeaderInfo, searchOutput) {
            this.searchOutput = searchOutput;
            this.tableHeaderInfo = tableHeaderInfo;
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

