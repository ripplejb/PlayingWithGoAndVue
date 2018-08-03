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

var app = new Vue({
    el: '#app',
    delimiters: ['${', '}'],
    data: {
        searchInput: '',
        searchOutput: {},
        tableHeaderInfo: {}
    },
    methods: {
        createDataCellDef(results, columnWidths) {
            var datainfo = [];
            var i = 0;
            for(var prop in results) {
               datainfo.push({value: results[prop], widthPercent: columnWidths[i++], cssClassName: 'mydatacell'})
            }
            return datainfo;
        },
        createHeaderDef(headers, columnWidths) {
            var headerInfo = [];

            for(var i=0; i < headers.length; i++) {
                headerInfo.push({value:headers[i], widthPercent: columnWidths[i], cssClassName:'myheadercell'})
            }

            return headerInfo;
        },
        sendData() {
            this.$http.get('/search?searchInput=' + this.searchInput).then((response) => {
                this.searchOutput = JSON.parse(response.bodyText);
                this.tableHeaderInfo = this.createHeaderDef(this.searchOutput.Headers, this.searchOutput.ColumnWidths)
                for(let ind in this.searchOutput.Results) {
                    this.searchOutput.Results[ind]['celldef'] = this.createDataCellDef(this.searchOutput.Results[ind], this.searchOutput.ColumnWidths);
                }
            });
        },
        addRowToDB(res) {
            if (res == undefined) return;
            this.$http.post('/books/add', res).then((response) => {
                alert("added the record");
            });
        }
    }
});

