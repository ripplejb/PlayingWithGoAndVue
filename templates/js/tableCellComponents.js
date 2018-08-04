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

function createDataCellDef(results, columnWidths) {
    var datainfo = [];
    var i = 0;
    for(var prop in results) {
        datainfo.push({value: results[prop], widthPercent: columnWidths[i++], cssClassName: 'mydatacell'})
    }
    return datainfo;
}

function createHeaderDef(headers, columnWidths) {
    var headerInfo = [];

    for(var i=0; i < headers.length; i++) {
        headerInfo.push({value:headers[i], widthPercent: columnWidths[i], cssClassName:'myheadercell'})
    }

    return headerInfo;
}
