CREATE TRIGGER reduce_queued_events
AFTER INSERT ON subscriptions.queue
FOR EACH ROW
WHEN ((NEW).allow_reduce)
EXECUTE FUNCTION subscriptions.reduce_queued_events();
