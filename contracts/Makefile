create-task:
	cast send $(AVS_ADDRESS) \
	  "submitTask(uint32,bytes)" \
	  $(EXECUTOR_OPERATOR_SET_ID) $(PAYLOAD) \
	  --private-key $(PRIVATE_KEY_APP) \
	  --rpc-url $(RPC_URL) \
	  --legacy
