var app = new Vue({
    el: '#app',
    delimiters: ['${', '}'],
    data: {
        searchInput: '',
        searchOutput: {},
        tableHeaderInfo: {}
    },
    methods: {
        sendData() {
            this.$http.get('/search?searchInput=' + this.searchInput).then((response) => {
                this.searchOutput = JSON.parse(response.bodyText);
                this.tableHeaderInfo = createHeaderDef(this.searchOutput.Headers, this.searchOutput.ColumnWidths)
                for(let ind in this.searchOutput.Results) {
                    this.searchOutput.Results[ind]['celldef'] = createDataCellDef(this.searchOutput.Results[ind], this.searchOutput.ColumnWidths);
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

