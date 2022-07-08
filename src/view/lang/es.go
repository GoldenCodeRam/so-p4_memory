package lang

const (
    IS_BLOCKED = "¿Se bloquea?"
    CREATE = "Crear"
    CREATE_PROCESS = "Crear proceso"
    CREATE_PARTITION = "Crear partición"
    NAME = "Nombre"
    TIME = "Tiempo"
    PROCESSES = "Procesos"
    PARTITIONS = "Particiones"
    PARTITION_NUMBER = "Número de partición"
    PARTITION_SIZE = "Tamaño de partición"

    // States
    INPUT_LIST = "Lista de entradas"
    READY_LIST = "Lista de listos"
    RUNNING_LIST = "Lista de en-ejecución"
    BLOCKED_LIST = "Lista de bloqueados"
    FINISHED_LIST = "Lista de finalizados"

    // Transitions
    READY_TO_RUNNING = "Listos a en-ejecución"
    RUNNING_TO_READY = "Tiempo expirado"
    RUNNING_TO_BLOCKED = "En-ejecución a bloqueado"
    BLOCKED_TO_READY = "E/S terminado"
    RUNNING_TO_FINISHED = "En-ejecución a terminado"


    ERROR_INVALID_NUMBER = "Error: número inválido."
    ERROR_EMPTY_TEXT = "Error: texto vacío."
    ERROR_COULD_NOT_ADD_PROCESS = "Error: no se pudo añadir el proceso."
)
