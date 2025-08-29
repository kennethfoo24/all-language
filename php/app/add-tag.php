<?php

// Getting the active span to add the tag to the active span
$span = \DDTrace\active_span();
if ($span) {
	$span->meta['status'] = 'active_span';
}

echo 'Adding a tag';
?>