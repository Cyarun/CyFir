#\!/bin/bash
# Final i18n fixes for remaining Velociraptor references

echo "=== Final i18n Translation Fixes ==="

# Spanish - fix remaining translations
sed -i 's/"Velociraptor Binary":"Ejecutable de Velociraptor"/"Velociraptor Binary":"Ejecutable de CyFir"/g' gui/velociraptor/src/components/i8n/es.jsx
sed -i 's/que se necesite la herramienta, Velociraptor la descargará/que se necesite la herramienta, CyFir la descargará/g' gui/velociraptor/src/components/i8n/es.jsx
sed -i 's/próxima actualización del servidor Velociraptor/próxima actualización del servidor CyFir/g' gui/velociraptor/src/components/i8n/es.jsx

# Portuguese - fix remaining 
sed -i 's/"Velociraptor Binary":"Binário do Velociraptor"/"Velociraptor Binary":"Binário do CyFir"/g' gui/velociraptor/src/components/i8n/por.jsx
sed -i 's/a ferramenta for necessária, o Velociraptor baixará/a ferramenta for necessária, o CyFir baixará/g' gui/velociraptor/src/components/i8n/por.jsx
sed -i 's/atualização do servidor Velociraptor/atualização do servidor CyFir/g' gui/velociraptor/src/components/i8n/por.jsx

# Vietnamese - fix remaining
sed -i 's/"Velociraptor Binary":"Tệp thực thi Velociraptor"/"Velociraptor Binary":"Tệp thực thi CyFir"/g' gui/velociraptor/src/components/i8n/vi.jsx

# Japanese - fix remaining
sed -i 's/"Velociraptor Binary":"Velociraptorバイナリ"/"Velociraptor Binary":"CyFirバイナリ"/g' gui/velociraptor/src/components/i8n/jp.jsx

# German - check for any missed ones
sed -i 's/Velociraptor-Server/CyFir-Server/g' gui/velociraptor/src/components/i8n/de.jsx

# French - check for any missed ones
sed -i 's/"Velociraptor Binary":"Binaire Velociraptor"/"Velociraptor Binary":"Binaire CyFir"/g' gui/velociraptor/src/components/i8n/fr.jsx

echo "=== Final Fixes Complete ==="
