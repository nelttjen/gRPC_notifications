from abc import ABC, abstractmethod


class BaseRPCRequestClass(ABC):
    @abstractmethod
    def perform_request(self) -> 'RPCResponse':
        pass

    @abstractmethod
    def process_response(self, response: 'RPCResponse') -> bool:
        pass