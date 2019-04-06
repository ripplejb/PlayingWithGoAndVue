
function createDataCellDef(results, columnWidths) {
    var datainfo = [];
    var i = 0;
    for(var prop in results) {
        if (i === columnWidths.length) break;
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

function getDividends(app) {
    app.$http.get('/Dividend?symbol=' + app.searchInput).then((response) => {
        app.searchOutput = JSON.parse(response.bodyText);
        app.tableHeaderInfo = createHeaderDef(app.searchOutput.Headers, app.searchOutput.ColumnWidths);
        Results = [];
        app.searchOutput['Results'] = Results;
        for(let ind in app.searchOutput.Dividends) {
            app.searchOutput.Results.push(app.searchOutput.Dividends[ind]);
            app.searchOutput.Results[ind]['celldef'] = createDataCellDef(app.searchOutput.Results[ind], app.searchOutput.ColumnWidths);
        }
        app.$emit('data-received', app.tableHeaderInfo, app.searchOutput);
    });
}

function getEarnings(app) {
    app.$http.get('/Earning?symbol=' + app.searchInput).then((response) => {
        app.searchOutput = JSON.parse(response.bodyText);
        app.tableHeaderInfo = createHeaderDef(app.searchOutput.Headers, app.searchOutput.ColumnWidths);
        Results = [];
        app.searchOutput['Results'] = Results;
        for(let ind in app.searchOutput.EarningsData.earnings) {
            app.searchOutput.Results.push(app.searchOutput.EarningsData.earnings[ind]);
            app.searchOutput.Results[ind]['celldef'] = createDataCellDef(app.searchOutput.EarningsData.earnings[ind], app.searchOutput.ColumnWidths);
        }
        app.$emit('data-received', app.tableHeaderInfo, app.searchOutput);
    });
}

