<?php
// This is the function that will lead to us add the tag
function tracedFunction (){
	sleep(1);
}

// Here we tell the tracer that if the function `addTag` is called then we add the tag to the current span
\DDTrace\trace_function(
	'tracedFunction',
	function (\DDTrace\SpanData $span, array $args, $retval, $exception) {

    	$span->name = 'Manual';
		$span->resource = 'create.span';
        $span->service = 'php_manual';
		// Adding a tag
		$span->meta['state'] = 'manual';

	}
);
	
// We call the function, it add a span with the second_tag
tracedFunction();

echo 'Adding a span';

?>