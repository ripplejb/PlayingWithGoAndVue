Vue.component('datacell', {
    delimiters: ['${', '}'],
    props: {
        widthPercent: Number,
        value: String,
        cssClassName: String
    },
    template: '<td :width="widthPercent" :class="cssClassName"> ${value}</td>'
});

Vue.component('headercell', {
    delimiters: ['${', '}'],
    props: {
        widthPercent: Number,
        value: String,
        cssClassName: String
    },
    template: '<th :width="widthPercent" :class="cssClassName"> ${value}</th>'
});

headerinfo =
[
    {
        widthPercent: 40,
        value: 'Title',
        cssClassName: 'myheadercell'
    },
    {
        widthPercent: 30,
        value: 'Author',
        cssClassName: 'myheadercell'
    },
    {
        widthPercent: 10,
        value: 'Year',
        cssClassName: 'myheadercell'
    },
    {
        widthPercent: 20,
        value: 'ID',
        cssClassName: 'myheadercell'
    }
];


var app = new Vue({
    el: '#app',
    delimiters: ['${', '}'],
    data: {
        searchInput: '',
        searchOutput: {},
        tableHeaderInfo: headerinfo
    },
    methods: {
        createDataCellDef(res) {
            datainfo =
                [
                    {
                        widthPercent: 40,
                        value: res.Title,
                        cssClassName: 'mydatacell'
                    },
                    {
                        widthPercent: 30,
                        value: res.Author,
                        cssClassName: 'mydatacell'
                    },
                    {
                        widthPercent: 10,
                        value: res.Year,
                        cssClassName: 'mydatacell'
                    },
                    {
                        widthPercent: 20,
                        value: res.ID,
                        cssClassName: 'mydatacell'
                    }
                ];
            return datainfo;
        },
        sendData() {
            this.$http.get('/search?searchInput=' + this.searchInput).then((response) => {
                this.searchOutput = JSON.parse(response.bodyText);
                for(let ind in this.searchOutput) {
                    this.searchOutput[ind]['celldef'] = this.createDataCellDef(this.searchOutput[ind]);
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

