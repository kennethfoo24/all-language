<?php

function doRiskyThing() {
    sleep(1);
    throw new Exception('Oops!');
}

\DDTrace\trace_function(
    'doRiskyThing',
    function() {
        // Span will be flagged as erroneous and have
        // the stack trace and exception message attached as tags
});

doRiskyThing();

echo 'Throwing an exception';

?>