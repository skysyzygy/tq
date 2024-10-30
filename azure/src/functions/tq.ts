import { app, HttpRequest, HttpResponseInit, InvocationContext } from "@azure/functions";
import { spawnSync } from 'node:child_process';

export async function tq(request: HttpRequest, context: InvocationContext): Promise<HttpResponseInit> {
    context.log(`Http function tq processed request for url "${request.url}"`);
    const {verb, object, variant, ...query} = request.params;
    var flag = "";
    if (variant != null) {
        flag = "--"+variant;
    } 

    var tq = spawnSync('bin/tq', ["-c", "--no-highlight", verb, object, flag], 
      {
        encoding: 'utf8', 
        input: JSON.stringify(query)
      });
    
    console.log(tq.stderr)

    if (tq.error) {
        console.error(tq.error.message);
    } else {
        return {jsonBody: JSON.parse(tq.stdout)};
    }
};

app.http('tq', {
    methods: ['GET', 'POST'],
    authLevel: 'anonymous',
    route: 'tq/{verb:alpha}/{object:alpha}/{variant:alpha?}',
    handler: tq
});
