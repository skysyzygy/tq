import { app, HttpRequest, HttpResponseInit, InvocationContext } from "@azure/functions";
import { spawnSync } from 'node:child_process';
import _ from 'lodash';

export async function tq(request: HttpRequest, context: InvocationContext): Promise<HttpResponseInit> {
    context.log(`tq processing ${request.method} request for url "${request.url}"`);
    const {object, variant, ...query} = request.params;
    const verb = request.method;
    var flag = "";
    if (variant != null) {
        flag = "--"+variant;
    }
     
    var tq = spawnSync('bin/tq', ["-c", "--no-highlight", verb, object, flag], 
      {
        encoding: 'utf8', 
        input: JSON.stringify(query),
        env: _.extend(process.env,{"TQ_LOGIN": request.headers.get("TQ_LOGIN")}),
        timeout: 30000
      });

    if (tq.status != 0) {
      return {
        status: 400,
        jsonBody: {
          status: tq.status,
          message: tq.stderr
      }}
    } else {
      return {jsonBody: JSON.parse(tq.stdout)};
    }

};

app.http('tq', {
    methods: ['GET', 'PUT', 'POST'],
    authLevel: 'anonymous',
    route: 'tq/{object:alpha}/{variant:alpha?}',
    handler: tq
});
